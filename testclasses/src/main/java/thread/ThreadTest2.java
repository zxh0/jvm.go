package thread;

public class ThreadTest2 extends Thread {
    
    public static void main(String[] args) {
        new ThreadTest2().start();
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
