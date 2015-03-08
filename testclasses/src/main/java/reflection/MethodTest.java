package reflection;

import java.lang.reflect.Method;

public class MethodTest implements Runnable {
    
    public static void main(String[] args) throws Exception {
        Method run = Runnable.class.getMethod("run");
        run.invoke(new MethodTest());
    }

    @Override
    public void run() {
        System.out.println("OK!");
    }
    
}
