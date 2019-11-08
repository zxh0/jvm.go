package jvm.instructions;

import static helper.MyAssert.*;

public class RefInvokeSpecialTest implements Runnable {

    static class B {
        int foo() { return 100; }
        int bar() { return 200; }
    }
    static class C1 extends B {
        int foo() { return super.foo(); } // invokespecial B.foo
    }
    static class C2 extends C1 {
        int foo() { return super.foo(); } // invokespecial C1.foo
        int bar() { return super.bar(); } // invokespecial C1.foo
    }

    public static void main(String[] args) {
        new RefInvokeSpecialTest().run();
    }

    @Override
    public void run() {
        assertEquals(100, new C2().foo());
        assertEquals(200, new C2().bar());
    }

}
