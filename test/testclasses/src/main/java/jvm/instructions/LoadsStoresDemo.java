package jvm.instructions;

public class LoadsStoresDemo {

    private static void iLoadStore(int a0, int a1, int a2, int a3, int a4) {
        a0 = a0; // iload_0, istore_0
        a1 = a1; // iload_1, istore_1
        a2 = a2; // iload_2, istore_2
        a3 = a3; // iload_3, istore_3
        a4 = a4; // iload,   istore
    }

    private static void lLoadStore(long a0, long a2, long a4) {
        a0 = a0; // lload_0, lstore_0
        a2 = a2; // lload_2, lstore_2
        a4 = a4; // lload,   lstore
    }

    private static void fLoadStore(float a0, float a1, float a2, float a3, float a4) {
        a0 = a0; // fload_0, fstore_0
        a1 = a1; // fload_1, fstore_1
        a2 = a2; // fload_2, fstore_2
        a3 = a3; // fload_3, fstore_3
        a4 = a4; // fload,   fstore
    }

    private static void dLoadStore(double a0, double a2, double a4) {
        a0 = a0; // dload_0, dstore_0
        a2 = a2; // dload_2, dstore_2
        a4 = a4; // dload,   dstore
    }

    private static void aLoadStore(Object a0, Object a1, Object a2, Object a3, Object a4) {
        a0 = a0; // aload_0, astore_0
        a1 = a1; // aload_1, astore_1
        a2 = a2; // aload_2, astore_2
        a3 = a3; // aload_3, astore_3
        a4 = a4; // aload,   astore
    }

    private static void xaLoadStore(int[] ia, long[] la, float[] fa, double[] da,
                                    Object[] aa, byte[] ba, char[] ca, short[] sa) {
        ia[0] = ia[0]; // iaload, iastore
        la[1] = la[1]; // laload, lastore
        fa[2] = fa[2]; // faload, fastore
        da[3] = da[3]; // daload, dastore
        aa[4] = aa[4]; // aaload, aastore
        ba[5] = ba[5]; // baload, bastore
        ca[6] = ca[6]; // caload, castore
        sa[7] = sa[7]; // saload, sastore
    }

}
