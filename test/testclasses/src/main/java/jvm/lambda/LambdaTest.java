package jvm.lambda;

import java.util.concurrent.Callable;
import static helper.MyAssert.*;

public class LambdaTest implements Runnable {

    public static void main(String[] args) {
        new LambdaTest().run();
    }

    @Override
    public void run() {
        Callable<String> r = () -> "hello";
        try {
            String x = r.call();
            assertEquals("hello", x);
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }

}
