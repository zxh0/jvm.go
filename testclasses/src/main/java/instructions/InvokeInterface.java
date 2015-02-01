package instructions;

public class InvokeInterface implements Runnable{
    
    public static void main(String[] args) {
        new InvokeInterface().run();
    }

    @Override
    public void run() {
        System.out.println("run!");
    }
    
}
