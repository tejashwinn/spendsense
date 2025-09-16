app: SplitwiseClone
language: golang
style: microservice
version: 1.0.0

entities:
  - name: User
    fields:
      - { name: id, type: uuid, primary: true }
      - { name: name, type: string, required: true }
      - { name: email, type: string, unique: true, required: true }
      - { name: password_hash, type: string, required: true }
      - { name: preferred_currency, type: string, default: "USD" }
      - { name: created_at, type: datetime, default: now }

  - name: Group
    fields:
      - { name: id, type: uuid, primary: true }
      - { name: name, type: string, required: true }
      - { name: currency, type: string, default: "USD" }
      - { name: created_by, type: uuid, ref: User.id }
      - { name: created_at, type: datetime, default: now }

  - name: GroupMember
    fields:
      - { name: group_id, type: uuid, ref: Group.id }
      - { name: user_id, type: uuid, ref: User.id }
      - { name: role, type: enum, values: ["member","admin"], default: "member" }
    primary_key: [group_id, user_id]

  - name: Expense
    fields:
      - { name: id, type: uuid, primary: true }
      - { name: group_id, type: uuid, ref: Group.id }
      - { name: description, type: string }
      - { name: amount, type: decimal(12,2), required: true }
      - { name: paid_by, type: uuid, ref: User.id }
      - { name: date, type: date, default: today }
      - { name: category, type: string }
      - { name: created_at, type: datetime, default: now }

  - name: ExpenseSplit
    fields:
      - { name: expense_id, type: uuid, ref: Expense.id }
      - { name: user_id, type: uuid, ref: User.id }
      - { name: amount, type: decimal(12,2) }
    primary_key: [expense_id, user_id]

api:
  style: REST
  auth: JWT
  routes:
    - { method: POST, path: /users, action: create_user }
    - { method: POST, path: /auth/login, action: login }
    - { method: GET, path: /users/{id}, action: get_user, auth: true }
    - { method: POST, path: /groups, action: create_group, auth: true }
    - { method: POST, path: /groups/{id}/members, action: add_member, auth: true }
    - { method: POST, path: /groups/{id}/expenses, action: add_expense, auth: true }
    - { method: GET, path: /groups/{id}/balances, action: get_group_balances, auth: true }
    - { method: POST, path: /groups/{id}/settle, action: settle_group, auth: true }

logic:
  balances:
    - input: expenses + splits
    - compute: net_balance per user
    - minimize: transactions (graph simplification)

non_functional:
  performance:
    latency_target_ms: 200
    concurrency_target: 10000
  security:
    password_hashing: bcrypt
    enforce_https: true
  observability:
    metrics: prometheus
    logs: json
    healthcheck: /health
  deployment:
    container: docker
    orchestration: kubernetes
    ci_cd: go test + lint + build + deploy

integrations:
  storage:
    receipts: s3/minio
  currency_rates:
    provider: exchangerate.host
  notifications:
    email: smtp
    push: optional
