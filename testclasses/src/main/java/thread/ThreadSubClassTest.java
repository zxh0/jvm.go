package thread;

public class ThreadSubClassTest extends Thread {
    
    public static void main(String[] args) {
        new ThreadSubClassTest().start();
        for (int i = 0; i < 100; i++) {
            System.out.println("main:" + i);
        }
    }

    @Override
    public void run() {
        for (int i = 0; i < 100; i++) {
            System.out.println("run:" + i);
        }
    }
    
}
