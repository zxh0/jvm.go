package main;

import jvm.instructions.LookupSwitchTest;
import jvm.instructions.MathTest;
import jvm.instructions.TableSwitchTest;

public class AllTests {

    public static void main(String[] args) {
        runTest(new MathTest());
        runTest(new TableSwitchTest());
        runTest(new LookupSwitchTest());
        System.out.println("OK!");
    }

    private static void runTest(Runnable test) {
        test.run();
        System.out.print("[ok] ");
        System.out.println(test.getClass().getName());
    }

}
