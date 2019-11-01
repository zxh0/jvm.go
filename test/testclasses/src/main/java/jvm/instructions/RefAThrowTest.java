package jvm.instructions;

import static helper.MyAssert.*;

public class RefAThrowTest implements Runnable {

    public static void main(String[] args) {
        new RefAThrowTest().run();
    }

    @Override
    public void run() {
        try {
            foo();
            fail();
        } catch (RuntimeException e) {
            assertEquals("!", e.getMessage());
        }
    }

    private void foo() {
        bar();
    }

    private void bar() {
        throw new RuntimeException("!");
    }

}
