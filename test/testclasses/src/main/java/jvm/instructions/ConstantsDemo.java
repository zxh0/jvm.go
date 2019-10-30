package jvm.instructions;

public class ConstantsDemo {

    public static void main(String[] args) {
        Object a = null; // aconst_null
        int im1 = -1;    // iconst_m1
        int i0 = 0;      // i_const0
        int i1 = 1;      // i_const1
        int i2 = 2;      // i_const2
        int i3 = 3;      // i_const3
        int i4 = 4;      // i_const4
        int i5 = 5;      // i_const5
        long l0 = 0;     // l_const0
        long l1 = 1;     // l_const1
        float f0 = 0;    // f_const0
        float f1 = 1;    // f_const1
        float f2 = 1;    // f_const2
        double d0 = 0;   // d_const0
        double d1 = 1;   // d_const1
        byte b1 = 100;   // bipush
        short s1 = 1000; // sipush
        String s = "?";  // ldc
    }

}
