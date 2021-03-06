## ServerlessFramework + Golang Practice
### 実装内容
- SQSでキューを受け取り、外部のAPIを叩き、結果をS3に保存するLambda関数
    - [x] SQSでキューを受け取る
    - [x] 外部のAPIを叩く
        - https://api.thecatapi.com/v1/images/search を叩いた結果のJSONを保存
    - [x] S3に保存
    - serverless.yml (CloudFormation)の設定
        - [x] Lambda関数設定
        - [x] SQS設定
- S3のファイル作成イベントを受け取り、ファイルを取得し、JSONをRDSに保存するLambda関数
    - [x] S3ファイル作成イベントを受け取る
    - [x] ファイルを取得
    - [x] JSONをパースしてRDSに保存
        - [x] RDS PROXY経由で接続
    - serverless.yml (CloudFormation)の設定
        - [x] Lambda関数設定
        - [x] VPC,Subnet設定
        - [x] Lambda関数をVPC内に
        - [x] RDS設定
        - [x] RDS Proxy設定
        - [x] RDSにアクセスするための踏み台EC2
- httpのGETリクエストを受け取り、RDSを検索した結果を返すLambda関数
    - [x] httpのGETリクエストを受け取る
    - [x] RDSを検索した結果を返す
- graphqlで実装された、RDSを検索した結果を返すLambda関数
    - [x] gqlgenを使う
    - [ ] appSyncを使う

