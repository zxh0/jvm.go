package jvmgo.thread;

public class SleepTest {
    
    public static void main(String[] args) throws InterruptedException {
        System.out.println("Before sleep!");
        Thread.sleep(5000);
        System.out.println("After sleep!");
    }
    
}
