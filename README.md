# プロジェクト概要
このプロジェクトは、月ごとの収入・支出・控除・貯金の進捗を一元管理できる家計簿アプリです。
本業・副業の収入、家賃や光熱費などの固定費、給与変動（残業や遅刻）、さらに貯金目標の達成状況をまとめて可視化できます。

ユーザーは毎月発生するお金の動きを簡単に登録でき、
「いまどれだけ貯金できているか」「欲しい物の金額まであとどれくらいか」
をすぐに確認できるよう設計されています。

バックエンドは Go + PostgreSQL を採用しており、
拡張性と読みやすさを重視した構成で開発を進めています。

## 技術スタック

### バックエンド
[![My Skills](https://skillicons.dev/icons?i=go)](https://skillicons.dev)

### フロントエンド
[![My Skills](https://skillicons.dev/icons?i=ts)](https://skillicons.dev)
[![My Skills](https://skillicons.dev/icons?i=react)](https://skillicons.dev)
[![My Skills](https://skillicons.dev/icons?i=tailwindcss)](https://skillicons.dev)

### インフラ・その他
[![My Skills](https://skillicons.dev/icons?i=docker,git,github,postgres,node)](https://skillicons.dev)


## 主な機能

- 月ごとの収入管理（本業・副業など）
- 支出管理（家賃・光熱費・保険・食費など）
- 給与変動（遅刻や欠勤・その他控除）
- 欲しいものに対しての貯金進捗の確認

## 環境変数

このアプリを起動するには、以下の環境変数を設定する必要があります。

| 変数名      | サンプル値         | 説明                       |
| -------- | ------------- | ------------------------ |
| `VOLUME` | `./api/.data` | PostgreSQL のデータ永続化ディレクトリ |
| `PORT` | `9090` | API サーバーの起動ポート |
| `POSTGRES_DATABASE` | `kakebo`      | データベース名                            |
| `POSTGRES_USER`     | `kakebo_user` | DB 接続ユーザー                          |
| `POSTGRES_PASSWORD` | `kakebo_pass` | DB 接続パスワード                         |
| `POSTGRES_PORT`     | `5432`        | PostgreSQL のポート                    |
| `POSTGRES_HOST`     | `postgres`    | DB ホスト（docker-compose の service 名） |
| `POSTGRES_SSLMODE`  | `disable`     | SSL の使用設定（ローカルでは disable）          |
| `GOOSE_DRIVER`        | `postgres`                                                                | Goose が使用するドライバ     |
| `GOOSE_DBSTRING`      | `postgres://kakebo_user:kakebo_pass@postgres:5432/kakebo?sslmode=disable` | DB 接続文字列            |
| `GOOSE_MIGRATION_DIR` | `tools/postgres/migrations`                                               | マイグレーションファイルのディレクトリ |
| `GOOSE_TABLE`         | `goose_migrations`                                                        | Goose の管理テーブル名      |
| `JWT_PRIVATE_KEY` | `cert/private.pem` | 署名用の秘密鍵パス |
| `JWT_PUBLIC_KEY`  | `cert/public.pem`  | 検証用の公開鍵パス |






