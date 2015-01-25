package jvmgo;

import java.io.OutputStream;

/**
 * To let HelloWord work!
 */
public class StdoutOutputStream extends OutputStream {

    @Override
    public native void write(int b);
    
}
