package main;

import jvm.instructions.*;
import stdlib.basic.StringTest;
import stdlib.basic.WrapperTest;
import stdlib.basic.cl.ClassLoaderTest;
import stdlib.basic.reflection.ClassTest;
import stdlib.basic.reflection.PrimitiveClassTest;

public class AllTests {

    public static void main(String[] args) {
        runTest(new MathTest());
        runTest(new ControlTest());
        runTest(new ComparisonsTest());
        runTest(new RefFieldTest());
        runTest(new RefInvokeTest());
        runTest(new RefArrayTest());
        runTest(new RefAThrowTest());
        runTest(new StringTest());
        runTest(new WrapperTest());
        runTest(new PrimitiveClassTest());
//        runTest(new ClassTest());
        runTest(new ClassLoaderTest());
        System.out.println("OK!");
    }

    private static void runTest(Runnable test) {
        test.run();
        System.out.print("[ok] ");
        System.out.println(test.getClass().getName());
    }

}
