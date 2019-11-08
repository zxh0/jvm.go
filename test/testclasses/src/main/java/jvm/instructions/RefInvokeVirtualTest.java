package jvm.instructions;

import static helper.MyAssert.*;

public class RefInvokeVirtualTest implements Runnable {

    interface I { default int foo() { return 100; } }
    static class C1 implements I {}
    static class C2 extends C1 {}
    static class C3 extends C2 {}

    public static void main(String[] args) {
        new RefInvokeVirtualTest().run();
    }

    @Override
    public void run() {
        assertEquals(1, new C1().foo()); // invokevirtual C1.foo
        assertEquals(1, new C2().foo()); // invokevirtual C2.foo
        assertEquals(1, new C3().foo()); // invokevirtual C3.foo
    }

}
