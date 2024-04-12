# Tinder配對系統說明文件

## 概述
這是一個基於 gin 實現的 HTTP 伺服器,用於支援類似 Tinder 的配對系統。該伺服器提供了三個主要的 API 功能
1. 添加新用戶並尋找潛在匹配對象。
2. 移除指定名稱用戶。
3. 查詢潛在匹配用戶列表。
## 檔案架構
```
.
├── README.md
├── config
│   ├── config.go
│   └── config.yml
├── di
│   └── di.go
├── doc
├── handler
│   └── match_handler.go
├── logger
│   ├── logger.go
│   └── zaplog.go
├── middleware
│   └── logger.go
├── model
│   └── model.go
├── server
│   ├── router.go
│   └── server.go
├── service
│   ├── match_service.go
│   └── match_service_test.go
├── .gitignore
├── Dockerfile
└── go.mod
```

### 專案結構說明
- `config`: 系統配置模組。
    - `config.go`: 用於讀取和解析配置文件。
    - `config.yml`: 系統配置文件,用於配置服務器的參數。
- `di`: 依賴注入模組,用於管理服務器的依賴關係。
  - `di.go`: 用於初始化服務器的依賴關係。
- `doc`: 說明文件目錄,用於存放系統的說明文件。
  - `README.md`: Tinder配對系統說明文件,用於描述系統的功能和使用方法。
- `handler`: 用於處理 API 請求，也負責處理參數驗證。
  - `match_handler.go`: 配對請求處理模組,用於處理配對相關的 API 請求。
- `logger`: 日誌模組,用於記錄系統運行日誌。
  - `logger.go`: 定義logger主要對外部的介面。在此取得系統使用的logger。
  - `zaplog.go`: Zap套件相關設定。
- `middleware`: 中間件模組,用於處理 HTTP 請求的中介層。
  - `logger.go`: 記錄執行結束的相關資訊，包含有無系統錯誤訊息。
- `model`: 數據模型模組,用於定義系統中的數據模型。
- `server`: 服務器模組,用於啟動和運行 HTTP 服務器。
  - `router.go`: 路由設置模組,用於設置 API 請求的路由。
  - `server.go`: 服務器啟動模組,用於啟動和運行 HTTP 服務器。
- `service`: 服務模組,用於實現系統的業務邏輯。
  - `match_service.go`: 配對服務模組,用於實現用戶配對的業務邏輯。
  - `match_service_test.go`: 配對服務測試模組,用於對配對服務進行單元測試。
- `.gitignore`: Git 忽略文件,用於配置 Git 忽略規則。
- `Dockerfile`: Docker 鏡像構建文件,用於構建 Docker 鏡像。
- `go.mod`: Go 模組文件,用於管理系統的依賴關係。
---
## 時間複雜度分析
1. AddSinglePersonAndMatch
  - 時間複雜度: O(n)
  - 說明: 該方法首先通過 isPersonExist 檢查新用戶是否已存在,時間複雜度為 O(n)。 
  然後,根據性別將新用戶添加到相應的列表中,並再次對列表進行排序,時間複雜度為 O(log n)。
  最後,調用 matchPerson 尋找潛在匹配。使用二分搜尋,時間複雜度也為 O(log n)。
  所以時間複雜度是O(n) + O(log n)＝O(n)。如果不需要檢查新用戶是否存在則可優化時間複雜度為O(log n)。
2. RemoveSinglePerson
- 時間複雜度: O(n)。
- 說明: 該方法首先在 males 列表中線性搜索指定姓名的用戶,如果找到則將其移除。如果未找到,則在 females 列表中繼續搜索。因此,在最壞情況下,需要確認兩個完整列表,時間複雜度為 O(n)。
3. QuerySinglePeople
- 時間複雜度: O(1) 到 O(n)
- 說明: 該方法根據性別選擇相應的候選用戶列表,並返回指定數量的用戶。如果請求的數量小於等於候選列表的長度,則只需返回列表的前 N 個元素,時間複雜度為 O(1)。但如果請求的數量大於候選列表的長度,則返回整個列表,時間複雜度為 O(n)。
---
## API說明文件
1. AddSinglePersonAndMatch
 - 說明: 添加一個新用戶到配對系統,並尋找潛在匹配對象。
 - HTTP方法 : `POST`
 - 路徑 : `/add`
 - 請求參數:
    - name（string,必須）: 用戶姓名
    - height: （int,必須）：用戶身高
    - gender: （string,必須）：用戶性別（male/female）
    - remainingDates (int,必須): 用戶剩餘約會次數
    - request範例:
    ```json
       {
           "name": "John",
           "height": 180,
           "gender": "male",
           "remainingDates": 2
       }
    ```
 - 回應參數
   - matchedPerson: 匹配的用戶資訊
       - name（string,必須）: 用戶姓名
       - height: （int,必須）：用戶身高
       - gender: （string,必須）：用戶性別（male/female）
       - remainingDates (int,必須): 用戶剩餘約會次數
       1. 成功配對 (HTTP 201 Created)
          ```json
          {
              "matchedPerson": 
              {
                  "name": "Lucy",
                  "height": 151,
                  "gender": "female",
                  "remainingDates": 2
              }
          }
          ```
     2. 未找到匹配 (HTTP 404 Not Found)
          ```json
          {
              "message": "No suitable match found"
          }
          ```
     3. 請求格式錯誤 (HTTP 400 Bad Request)
          ```json
          {
              "message": "invalid input"
          }
          ```
2. RemoveSinglePerson
 - 說明: 從配對系統中移除指定用戶。
 - HTTP方法 : `DELETE`
 - 路徑 : `/remove/:name`
 - 請求參數:
    - name (string, 必須): 要移除的用戶姓名
- 回應參數
  - 成功移除 (HTTP 204 No Content)
      ```json
        {
            "message" : "Single person removed"
        }
     ```
  - 用戶不存在 (HTTP 404 Not Found)
    ```json
      {
          "error": "Single person with name '{name}' not found"
      }
    ```
3. QuerySinglePeople
    - 說明: 查詢配對系統中的用戶。
    - HTTP方法 : `GET`
    - 路徑 : `/query?gender={string}&limit={int}`
    - 請求參數:
      - gender (string, 必須): 用戶性別（male/female）
      - limit (int, 必須): 返回的用戶數量
    - 回應參數:
      - 成功查詢 (HTTP 200 OK)
        - Array of matchedPerson
          - name（string,必須）: 用戶姓名
          - height（int,必須）：用戶身高
          - gender（string,必須）：用戶性別（male/female）
          - remainingDates(int,必須): 用戶剩餘約會次數
        ```json
           [
               {
                   "name": "Mike",
                   "height": 165,
                   "gender": "male",
                   "remainingDates": 2
               },
               {
                   "name": "Hank",
                   "height": 166,
                   "gender": "male",
                   "remainingDates": 2
               }
           ]
        ```
      - gender參數不合法(HTTP 400 Bad Request)
        ```json
        {
            "error": "gender is invalid"
        } 
        ```
      - 無效的limit參數
        ```json
        {
            "error": "limit invalid limit"
        }
        ```
---
## Getting started in docker
1.Clone the repository
```bash
git clone https://github.com/Chao-Shiun/BitoPro_interview_question.git
cd BitoPro_interview_question
```
2. Build dockerfile & create docker image
```bash
docker build -t tinder-matching-system .
```
3. Run docker container
```bash
docker run -p 8080:8080 tinder-matching-system
```
4. Stop and remove the container when you're done
```bash
docker stop <container_id>
docker rm <container_id>
```
---
## Todo-List
- [ ] 將`match_handler.go`驗證參數的邏輯抽取出來，讓程式碼更簡潔一致。
- [ ] 在service層也注入日誌功能，並且可以記錄錯誤來源與行數（caller）。這部份的設定會與middleware的logger.go有所不同。
- [ ] middleware的logger.go可以加入攔截更多不同情況。例如參數錯誤使用Warn log。系統錯誤除了記錄之外也可主動通知管理者。