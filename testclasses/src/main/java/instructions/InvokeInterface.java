package instructions;

public class InvokeInterface implements Runnable{
    
    public static void main(String[] args) {
        Runnable r = new InvokeInterface();
        r.run();
    }

    @Override
    public void run() {
        System.out.println("run!");
    }
    
}
