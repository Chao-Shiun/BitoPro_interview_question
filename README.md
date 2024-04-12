# Anwser 1

當這個程式執行時,會發生以下行為:

- 編譯時會發生`symbol too large (800000000000 bytes > 2000000000 bytes)`的錯誤，這是因為`[10e10]uint64{}`預期會建立約80GB左右的陣列，但Go編譯器預設的最大大小約為2GB。
- 程式模擬了兩個用戶(userA和userB)之間的轉帳操作。程式啟動後,會創建兩個goroutine,分別代表從userA轉帳到userB和從userB轉帳到userA。每個goroutine會執行10^10次轉帳操作,每次轉帳金額為1。

**下面是程式的執行過程:**

1. 程式啟動,創建userA和userB,初始餘額都為10^10。
2. 主goroutine啟動兩個子goroutine,每個goroutine執行10^10次轉帳操作。
3. 在每個goroutine中,對於每一次轉帳:
   - 呼叫transfer函式,傳入轉出方、轉入方和轉帳金額(1)。
   - 在transfer函式中,先鎖定轉出方,再鎖定轉入方。
   - 檢查轉出方餘額是否足夠,如果足夠,就執行轉帳,將1從轉出方餘額中扣除,加到轉入方餘額中。
   - 釋放轉入方和轉出方的鎖。
4. 兩個goroutine不斷地執行上述轉帳操作,直到每個goroutine都執行了10^10次。
5. 當兩個goroutine都完成後,主goroutine會繼續執行,程式結束。

**在這個過程中,可能發生以下問題**
1. deadlock：如果兩個goroutine同時執行到`transfer`函式,一個goroutine鎖定了`userA`,另一個goroutine鎖定了`userB`,然後它們都試圖獲取對方持有的鎖,就會發生死鎖。程式會卡住,無法繼續執行。

2. 效能問題:
   - 由於每次轉帳操作都需要獲取和釋放鎖,而且轉帳操作非常頻繁(每個goroutine執行10^10次),大部分時間都會消耗在鎖的競爭和管理上,而不是真正的轉帳操作。
   - 由於鎖的粒度較大(鎖住整個User物件),即使兩個goroutine訪問不同的User,也會發生鎖競爭。

**改善程式流程:**

- batchTransfer函式實現了從一個用戶轉帳到另一個用戶的邏輯。它接受轉出方from、轉入方to和轉帳金額amount作為參數。
- 在batchTransfer函式中,使用了一個特定的鎖定順序來避免死鎖。它總是先鎖定ID較小的用戶,然後再鎖定ID較大的用戶。這樣可以保證不同的goroutine在競爭鎖時,總是以相同的順序獲取鎖,從而避免死鎖。
- 在鎖定了兩個用戶後,batchTransfer函式檢查轉出方的餘額是否足夠,如果足夠,就執行轉帳操作,將指定金額從轉出方的餘額中扣除,並加到轉入方的餘額中。
- 在main函式中,創建了兩個初始餘額為10^10的用戶userA和userB。
- 然後啟動了兩個goroutine,分別代表從userA轉帳到userB和從userB轉帳到userA。每個goroutine執行10^10次轉帳操作,每次轉帳金額為1。
- 使用sync.WaitGroup來等待兩個goroutine完成。wg.Add(2)表示需要等待兩個goroutine,每個goroutine在退出前執行wg.Done()來通知WaitGroup它已經完成。
- 在兩個goroutine都完成後,主goroutine會執行wg.Wait()來等待它們的完成。




```go

import (
	"fmt"
	"sync"
)

type User struct {
	ID      uint64
	Balance uint64
	Lock    sync.Mutex
}

func batchTransfer(from *User, to *User, amount uint64) {
	// 總是先鎖定ID較小的用戶,避免deadlock
	if from.ID < to.ID {
		from.Lock.Lock()
		to.Lock.Lock()
	} else {
		to.Lock.Lock()
		from.Lock.Lock()
	}
	defer from.Lock.Unlock()
	defer to.Lock.Unlock()

	if from.Balance >= amount {
		from.Balance -= amount
		to.Balance += amount
	}
}

func main() {
	userA := User{
		ID:      1,
		Balance: 10e10,
	}
	userB := User{
		ID:      2,
		Balance: 10e10,
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := uint64(0); i < 10e10; i++ {
			batchTransfer(&userA, &userB, 1)
		}
	}()

	go func() {
		defer wg.Done()
		for i := uint64(0); i < 10e10; i++ {
			batchTransfer(&userB, &userA, 1)
		}
	}()

	wg.Wait()
	fmt.Printf("UserA Balance: %d, UserB Balance: %d\n", userA.Balance, userB.Balance)
}

```


# Answer 2
在Redis中儲存用戶最近購買的100個商品,可以採用以下設計:

- 對每個用戶,使用一個Hash結構儲存商品資訊,key為用戶ID,field為商品ID,value為商品詳情(JSON格式)。
- 同時用一個List結構記錄商品的購買時間順序,以便取出最近的100個商品。List的key為用戶ID,value為商品ID。
- 當新增商品時,在Hash中新增一個field,並將商品ID從List左端插入。如果List長度超過100,從右端彈出多餘的商品ID,並從Hash中刪除對應的field。

這樣可以保證List中始終包含該用戶最近購買的100個商品,而Hash提供商品的詳細信息。透過設置合理的key/field過期時間,可以控制記憶體用量。

# Answer 3
Rolling Update和Recreate是Kubernetes Deployment的兩種更新策略:

- Recreate策略是先刪除所有舊版本Pod,再建立新版本Pod,更新過程中服務會短暫中斷。
- Rolling Update策略是漸進式地用新版本Pod替換舊版本Pod,同一時間總是有可用的Pod,服務不中斷。

Rolling Update通常與readiness probe配合使用。readiness probe用於判斷新版本Pod是否已經就緒並可以接收流量。只有通過readiness probe的Pod才會被標記為可用,並逐步替換舊版本Pod。這樣確保更新過程平滑,服務質量不受影響。

# Answer 4
對於給定的SQL查詢,會同時用到user_id, created_at和status三個篩選條件。

- 索引A和B都包含這三列,順序略有不同。但對於給定的查詢順序(user_id AND created_at AND status),索引A的順序更匹配,因此效能會優於索引B。
- 索引C缺少status列,對status的篩選無法使用索引,因此效能最差。
  綜上,索引A的效能最好,索引C最差,B居中。使用複合索引時,列的順序要儘量匹配查詢條件的順序,才能發揮最大效能。

# Answer 5
在Kafka中,提升消費者端的效能主要有兩種方式:

1. 增加分區(Partition)數量,讓更多消費者實例並行消費。每個分區只能被每個消費者群組中的一個消費者讀取,因此提高分區數可以線性擴展消費能力。
2. 引入消費者群組(Consumer Group),每個群組獨立消費一份完整的資料。不同群組間可以並行消費資料,從而擴展消費能力。

但上述方案也有一些限制:

- 分區數事先確定,之後難以調整。分區過多也會增加管理成本。
- 消費者群組間是競爭關係,不適合某些需要協同的場景,如消息的全序消費。
- 消費狀態(offset)由Kafka維護,應用本身缺乏掌控。

應對這些限制的方法包括使用外部儲存管理offset、引入流處理引擎如Kafka Stream/Spark等。

# Answer 6
[Link Text](./doc/README.md)