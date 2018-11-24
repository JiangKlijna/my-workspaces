
import javax.imageio.ImageIO;
import java.awt.image.BufferedImage;
import java.io.File;
import java.io.IOException;
import java.util.concurrent.BlockingQueue;
import java.util.concurrent.LinkedBlockingDeque;

import static java.awt.image.BufferedImage.TYPE_INT_RGB;

public class Png {
 
    public static void main(String[] args) throws Exception {
        final BlockingQueue<Integer> queue = new LinkedBlockingDeque<Integer>();
        new Thread(() -> {
            //(256*256*256/2)/512
            BufferedImage img = new BufferedImage(4096, 4096, TYPE_INT_RGB);
            for (int i = 0; i < img.getWidth(); i++) {
                for (int j = 0; j < img.getHeight(); j++) {
                    try {
                        img.setRGB(i, j, queue.take());
                    } catch (InterruptedException e) {
                        e.printStackTrace();
                    }
                }
            }
            try {
                ImageIO.write(img, "png", new File("./test.png"));
            } catch (IOException e) {
                e.printStackTrace();
            }
        }).start();

        new Thread(() -> {
            for (int r = 0; r <= 255; r++) {
                for (int g = 0; g <= 255; g++) {
                    for (int b = 0; b <= 255; b++) {
                        try {
                            queue.put(Color(r, g, b));
                        } catch (InterruptedException e) {
                            e.printStackTrace();
                        }
                    }
                }
            }
        }).start();
    }


    public static int Color(int r, int g, int b) {
        return ((0xFF) << 24) |
                ((r & 0xFF) << 16) |
                ((g & 0xFF) << 8) |
                ((b & 0xFF) << 0);
    }
}
