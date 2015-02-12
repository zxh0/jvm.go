package thread;

public class ThreadTest implements Runnable {
    
    public static void main(String[] args) {
        new Thread(new ThreadTest()).start();
        for (int i = 0; i < 10; i++) {
            System.out.println("main");
        }
    }

    @Override
    public void run() {
        for (int i = 0; i < 10; i++) {
            System.out.println("run!");
        }
    }
    
}
