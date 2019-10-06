package gui.swing;

import javax.swing.JButton;
import javax.swing.JFrame;

public class SwingTest {
    
    public static void main(String[] args) {
        JFrame frame = new JFrame("Swing Test");
        frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        frame.getContentPane().add(new JButton("Click me!"));
        frame.setLocationRelativeTo(null);
        frame.pack();
        frame.setVisible(true);
    }
    
}
