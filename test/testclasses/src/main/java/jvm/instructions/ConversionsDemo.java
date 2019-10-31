package jvm.instructions;

public class ConversionsDemo {

    public static void main(String[] args) {
        int i = 0;
        long l = 0;
        float f = 0;
        double d = 0;
        byte b;
        char c;
        short s;

        l = i;        // i2l
        f = i;        // i2f
        d = i;        // i2d
        i = (int)l;   // l2i
        f = l;        // l2f
        d = l;        // l2d
        i = (int)f;   // f2i
        l = (long)f;  // f2l
        d = f;        // f2d
        i = (int)d;   // d2i
        l = (long)d;  // d2l
        f = (float)d; // d2f
        b = (byte)i;  // i2b
        c = (char)i;  // i2c
        s = (short)i; // i2s
    }

}
