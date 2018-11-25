
import javax.imageio.ImageIO;
import java.awt.image.BufferedImage;
import java.io.File;
import java.io.IOException;

import static java.awt.image.BufferedImage.TYPE_INT_RGB;

public abstract class PngBuilder implements Runnable {

    public final String filename;

    private PngBuilder(String filename) {
        this.filename = filename;
    }

    public static void main(String[] args) {
        new Thread(rgb).start();
        new Thread(rbg).start();
        new Thread(grb).start();
        new Thread(gbr).start();
        new Thread(brg).start();
        new Thread(bgr).start();
    }

    //(256*256*256/2)/512
    public static int LENGTH = 4096;

    @Override
    public void run() {
        BufferedImage img = new BufferedImage(LENGTH, LENGTH, TYPE_INT_RGB);
        for (int i = 0; i < LENGTH; i++) {
            for (int j = 0; j < LENGTH; j++) {
                img.setRGB(i, j, nextColor());
            }
        }
        try {
            ImageIO.write(img, "png", new File(filename));
			System.out.println("finish " + filename);
        } catch (IOException e) {
            e.printStackTrace();
        }
        System.gc();
    }

    private int i, j, k;

    public int nextColor() {
        if (k == 255) {
            if (j == 255) {
                if (i == 255) {
                    i = 0;
                    j = 0;
                    k = 0;
                } else {
                    i++;
                    j = 0;
                    k = 0;
                }
            } else {
                j++;
                k = 0;
            }
        } else {
            k++;
        }
        return builderColor(i, j, k);
    }

    public abstract int builderColor(int r, int g, int b);

    public static int getColor(int r, int g, int b) {
        return ((0xFF) << 24) |
                ((r & 0xFF) << 16) |
                ((g & 0xFF) << 8) |
                ((b & 0xFF) << 0);
    }

    public static final PngBuilder rgb = new PngBuilder("rgb.png") {
        @Override
        public int builderColor(int r, int g, int b) {
            return getColor(r, b, g);
        }
    };

    public static final PngBuilder grb = new PngBuilder("grb.png") {
        @Override
        public int builderColor(int r, int g, int b) {
            return getColor(g, r, b);
        }
    };
    public static final PngBuilder gbr = new PngBuilder("gbr.png") {
        @Override
        public int builderColor(int r, int g, int b) {
            return getColor(g, b, r);
        }
    };

    public static final PngBuilder rbg = new PngBuilder("rbg.png") {
        @Override
        public int builderColor(int r, int g, int b) {
            return getColor(r, g, b);
        }
    };

    public static final PngBuilder brg = new PngBuilder("brg.png") {
        @Override
        public int builderColor(int r, int g, int b) {
            return getColor(b, r, g);
        }
    };

    public static final PngBuilder bgr = new PngBuilder("bgr.png") {
        @Override
        public int builderColor(int r, int g, int b) {
            return getColor(b, g, r);
        }
    };

}
