package main
import "fmt"

func isPalindrome(num int) bool {
    original := num
    reversed := 0
    for num > 0 {
        reversed = reversed*10 + num%10
        num /= 10
    }
    return original == reversed
}

func palindromTerdekat(n int) {
    if isPalindrome(n) {
        fmt.Println(n, "adalah palindrom")
        return
    }

    lower := n - 1
    for lower >= 0 {
        if isPalindrome(lower) {
            break
        }
        lower--
    }

    higher := n + 1
    for {
        if isPalindrome(higher) {
            break
        }
        higher++
    }

    // Cari yang terdekat
    if n-lower <= higher-n {
        fmt.Println(lower)
    } else {
        fmt.Println(higher)
    }
}

func main() {
    var n int
    fmt.Print("Masukkan bilangan: ")
    fmt.Scan(&n)
    palindromTerdekat(n)
}