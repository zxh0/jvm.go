package gui.awt;

import java.awt.Button;
import java.awt.Frame;
import java.awt.event.WindowAdapter;
import java.awt.event.WindowEvent;

public class AwtTest extends WindowAdapter {
    
    public static void main(String[] args) {
        Frame frame = new Frame("AWT Test");
        Button button = new Button("Click me!");
        frame.add("Center", button);
        frame.addWindowListener(new AwtTest());
        frame.pack();
        frame.setVisible(true);
    }

    @Override
    public void windowClosing(WindowEvent e) {
        System.exit(0);
    }
    
}
