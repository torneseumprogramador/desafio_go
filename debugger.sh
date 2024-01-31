# go install github.com/go-delve/delve/cmd/dlv@latest

dlv debug main.go

# # Defina um breakpoint após a variável 'idade' ser inicializada
# (dlv) break main.go:8

# # Continue a execução até atingir o breakpoint
# (dlv) continue

# # vai para proxima linha do breakpoint
# (dlv) next

# # Agora tente imprimir a variável 'idade' novamente
# (dlv) print idade
