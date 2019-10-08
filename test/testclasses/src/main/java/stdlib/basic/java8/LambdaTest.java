package stdlib.basic.java8;

public class LambdaTest {
    
    public static void main(String[] args) {
        Runnable r = () -> {
            System.out.println("Hello, World!");
        };
        r.run();
    }
    
}
