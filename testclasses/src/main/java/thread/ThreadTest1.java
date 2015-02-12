package thread;

public class ThreadTest1 implements Runnable {
    
    public static void main(String[] args) {
        new Thread(new ThreadTest1()).start();
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
