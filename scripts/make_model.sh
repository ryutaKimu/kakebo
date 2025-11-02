#!/bin/bash
# 使い方: ./scripts/make_model.sh account

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

echo "✅ Created ${NAME} model"