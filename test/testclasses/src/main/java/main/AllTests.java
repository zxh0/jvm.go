package main;

import jvm.ClassInitTest;
import jvm.instructions.*;
import jvm.jsr292.LookupTest;
import stdlib.basic.StringTest;
import stdlib.basic.WrappersTest;
import stdlib.basic.cl.ClassLoaderTest;
import stdlib.basic.reflection.*;
import stdlib.x.unsafe.UnsafeMemoryTest;
import stdlib.x.unsafe.UnsafeObjectTest;

public class AllTests {

    public static void main(String[] args) {
        runTest(new MathTest());
        runTest(new ControlTest());
        runTest(new ComparisonsTest());
        runTest(new RefFieldTest());
        runTest(new RefInvokeTest());
        runTest(new RefArrayTest());
        runTest(new RefAThrowTest());
        runTest(new ClassInitTest());
        runTest(new StringTest());
        runTest(new WrappersTest());
        runTest(new ClassLoaderTest());
        runTest(new PrimitiveClassTest());
        runTest(new ArrayClassTest());
        runTest(new ClassTest());
        runTest(new ArrayTest());
        runTest(new UnsafeObjectTest());
        //runTest(new UnsafeMemoryTest());
        runTest(new LookupTest(args));
        System.out.println("OK!");
    }

    private static void runTest(Runnable test) {
        test.run();
        System.out.print("[ok] ");
        System.out.println(test.getClass().getName());
    }

}
