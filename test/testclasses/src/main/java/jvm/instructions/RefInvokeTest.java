package jvm.instructions;

public class RefInvokeTest implements Runnable {

    public static void main(String[] args) {
        Runnable r = new RefInvokeTest(); // invokespecial
        r.run(); // invokeinterface
        r.toString(); // invokevirtual
        staticMethod(); // invokestatic
    }

    private static void staticMethod() {
        // empty
    }

    private void privateMethod() {
        // empty
    }

    @Override
    public void run() {
        privateMethod(); // invokespecial ??
    }

    @Override
    public String toString() {
        return super.toString(); // invokespecial
    }

}
