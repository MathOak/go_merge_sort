# Go Merge Sort

Este projeto tem como objetivo aplicar e entender como executar uma função de merge sort em Go utilizando as Go routines.

## Como Rodar

1. **Clone o repositório:**
    ```sh
    git clone https://github.com/MathOak/go_merge_sort.git
    cd go_merge_sort
    ```

2. **Compile o código:**
    ```sh
    go build -o mergesort main.go
    ```

3. **Execute o programa:**
    ```sh
    ./mergesort
    ```

## Estrutura do Projeto

- `main.go`: Contém a implementação da função de merge sort utilizando Go routines.

## Motivo

O merge sort é um algoritmo de ordenação eficiente e estável que divide a lista em sub-listas menores, ordena essas sub-listas e, em seguida, as mescla para produzir a lista ordenada final. Utilizar Go routines permite que o processo de divisão e mesclagem seja executado de forma concorrente, aproveitando melhor os recursos do sistema e potencialmente melhorando o desempenho.

Este projeto foi criado para praticar a implementação de algoritmos concorrentes em Go e entender melhor como as Go routines podem ser utilizadas para melhorar a eficiência de algoritmos de ordenação.

