package main;

import jvm.instructions.*;

public class AllTests {

    public static void main(String[] args) {
        runTest(new MathTest());
        runTest(new ControlTest());
        runTest(new ComparisonsTest());
        runTest(new RefFieldTest());
        runTest(new RefInvokeTest());
        runTest(new RefArrayTest());
        runTest(new RefAThrowTest());
        System.out.println("OK!");
    }

    private static void runTest(Runnable test) {
        test.run();
        System.out.print("[ok] ");
        System.out.println(test.getClass().getName());
    }

}
