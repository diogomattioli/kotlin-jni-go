class Foo {
    init {
        System.loadLibrary("foo")
    }
    external fun add(a: Int, b: Int): Int
}

external fun mul(a: Int, b: Int): Int

fun main() {
    val c = Foo().add(1, 2)
    println(c)
    val d = mul(3, 4)
    println(d)
} 