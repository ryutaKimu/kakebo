#!/bin/bash
# 使い方: ./scripts/make_module.sh account

NAME=$1
CAPITALIZED=$(echo "$NAME" | sed 's/^\(.\)/\U\1/')

# Model
cat <<EOF >  internal/model/${NAME}.go
package model

type ${NAME} struct {
    ID   int64  \`json:"id"\`
    Name string \`json:"name"\`
}
EOF

# Controller
cat <<EOF > internal/controller/${NAME}_controller.go
package controller

type ${CAPITALIZED}Controller struct {
    service service.${CAPITALIZED}Service
}
EOF

# Service
cat <<EOF > internal/service/${NAME}_service.go
package service

type ${CAPITALIZED}Service struct {
    repo repository.${CAPITALIZED}Repository
}
EOF

# Service Interface
cat <<EOF > internal/service/interfaces/${NAME}_service.go
package interfaces

type ${CAPITALIZED}Service interface {
}
EOF

# Repository
cat <<EOF > internal/infra/postgre/${NAME}_repository.go
package repository

type ${CAPITALIZED}Repository struct{}
EOF

echo "✅ Created ${NAME} module (controller, model, service, interface repository)"
