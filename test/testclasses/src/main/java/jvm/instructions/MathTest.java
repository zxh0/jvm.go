package jvm.instructions;

import static helper.MyAssert.*;

public class MathTest implements Runnable {

    public static void main(String[] args) {
        new MathTest().run();
    }

    @Override
    public void run() {
        testIntOps();
        testLongOps();
        testFloatOps();
        testDoubleOps();
    }

    private static void testIntOps() {
        assertEquals(5, intOp(3, "+", 2));
        assertEquals(1, intOp(3, "-", 2));
        assertEquals(6, intOp(3, "*", 2));
        assertEquals(1, intOp(3, "/", 2));
        assertEquals(1, intOp(3, "%", 2));
        assertEquals(-2, intOp(0, "`-", 2));
        assertEquals(0xf0, intOp(0x0f, "<<", 4));
        assertEquals(-1, intOp(-1, ">>", 4));
        assertEquals(0xff, intOp(-1, ">>>", 24));
        assertEquals(0x00, intOp(0xf0, "&", 0x0f));
        assertEquals(0xff, intOp(0xf0, "|", 0x0f));
        assertEquals(0xf0, intOp(0xff, "^", 0x0f));
        assertEquals(4, intOp(0, "++", 3));
    }

    private static void testLongOps() {
        assertEquals(5L, longOp(3L, "+", 2L));
        assertEquals(1L, longOp(3L, "-", 2L));
        assertEquals(6L, longOp(3L, "*", 2L));
        assertEquals(1L, longOp(3L, "/", 2L));
        assertEquals(1L, longOp(3L, "%", 2L));
        assertEquals(-2L, longOp(0L, "`-", 2L));
        assertEquals(0xf0L, longOp(0x0fL, "<<", 4));
        assertEquals(-1L, longOp(-1L, ">>", 4));
        assertEquals(0xffL, longOp(-1L, ">>>", 56));
        assertEquals(0x00L, longOp(0xf0L, "&", 0x0fL));
        assertEquals(0xffL, longOp(0xf0L, "|", 0x0fL));
        assertEquals(0xf0L, longOp(0xffL, "^", 0x0fL));
    }

    private static void testFloatOps() {
        assertEquals(5.0f, floatOp(3f, "+", 2f));
        assertEquals(1.0f, floatOp(3f, "-", 2f));
        assertEquals(6.0f, floatOp(3f, "*", 2f));
        assertEquals(1.5f, floatOp(3f, "/", 2f));
        assertEquals(1.0f, floatOp(3f, "%", 2f));
        assertEquals(-2f, floatOp(0f, "`-", 2f));
    }

    private static void testDoubleOps() {
        assertEquals(5.0, doubleOp(3.0, "+", 2.0));
        assertEquals(1.0, doubleOp(3.0, "-", 2.0));
        assertEquals(6.0, doubleOp(3.0, "*", 2.0));
        assertEquals(1.5, doubleOp(3.0, "/", 2.0));
        assertEquals(1.0, doubleOp(3.0, "%", 2.0));
        assertEquals(-2.0, doubleOp(0.0, "`-", 2.0));
    }

    private static int intOp(int a, String op, int b) {
        switch (op) {
            case "+":   return a + b;   // iadd
            case "-":   return a - b;   // isub
            case "*":   return a * b;   // imul
            case "/":   return a / b;   // idiv
            case "%":   return a % b;   // irem
            case "`-":  return -b;      // ineg
            case "<<":  return a << b;  // ishl
            case ">>":  return a >> b;  // ishr
            case ">>>": return a >>> b; // iushr
            case "&":   return a & b;   // iand
            case "|":   return a | b;   // ior
            case "^":   return a ^ b;   // ixor
            case "++":  return ++b;     // iinc
            case "~":   return ~b;      //
            default:    return 0;
        }
    }

    private static long longOp(long a, String op, long b) {
        switch (op) {
            case "+":   return a + b;   // ladd
            case "-":   return a - b;   // lsub
            case "*":   return a * b;   // lmul
            case "/":   return a / b;   // ldiv
            case "%":   return a % b;   // lrem
            case "`-":  return -b;      // lneg
            case "<<":  return a << b;  // lshl
            case ">>":  return a >> b;  // lshr
            case ">>>": return a >>> b; // lushr
            case "&":   return a & b;   // land
            case "|":   return a | b;   // lor
            case "^":   return a ^ b;   // lxor
            default:    return 0;
        }
    }

    private static float floatOp(float a, String op, float b) {
        switch (op) {
            case "+":   return a + b;   // fadd
            case "-":   return a - b;   // fsub
            case "*":   return a * b;   // fmul
            case "/":   return a / b;   // fdiv
            case "%":   return a % b;   // frem
            case "`-":  return -b;      // fneg
            default:    return 0;
        }
    }

    private static double doubleOp(double a, String op, double b) {
        switch (op) {
            case "+":   return a + b;   // fadd
            case "-":   return a - b;   // fsub
            case "*":   return a * b;   // fmul
            case "/":   return a / b;   // fdiv
            case "%":   return a % b;   // frem
            case "`-":  return -b;      // fneg
            default:    return 0;
        }
    }

}
