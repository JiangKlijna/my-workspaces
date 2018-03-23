
import java.io.*;
import java.util.LinkedList;
import java.util.List;
import java.util.function.Consumer;
import java.util.jar.JarEntry;
import java.util.jar.JarInputStream;
import java.util.jar.JarOutputStream;

public final class Compression {
	private static final String VerboseCommand = "java -verbose";
	private static final ClassLoader CL = getClassLoader();
	private static final String ManifestInfo = getManifestInfo();

	public static void main(String[] args) {
		if (args.length == 0) {
			System.out.println("java Compression xxx.jar");
			return;
		}
		File appFile = new File(args[0]);
		try {
			new Compression(appFile).compress();
		} catch (IOException e) {
			System.out.println(e.getMessage());
		}
	}

	private static ClassLoader getClassLoader() {
		ClassLoader classloader = Thread.currentThread().getContextClassLoader();
		if (classloader == null) classloader = Compression.class.getClassLoader();
		if (classloader == null) classloader = ClassLoader.getSystemClassLoader();
		return classloader;
	}

	private static String getManifestInfo() {
		final String JAVA_VERSION = System.getProperty("java.version");
		final String JAVA_SPECIFICATION_VERSION = System.getProperty("java.specification.version");
		return "Manifest-Version: 1.0\n" +
				"Implementation-Vendor: Sun Microsystems, Inc.\n" +
				"Implementation-Title: Java Runtime Environment\n" +
				("Implementation-Version: " + JAVA_VERSION + "\n") +
				"Specification-Vendor: Sun Microsystems, Inc.\n" +
				("Created-By: " + JAVA_VERSION + " (Sun Microsystems Inc.)\n") +
				"Specification-Title: Java Platform API Specification\n" +
				("Specification-Version: " + JAVA_SPECIFICATION_VERSION);
	}

	private static void getLoadedClass(File jar, Consumer<String> action) throws IOException {
		Process process = Runtime.getRuntime().exec(VerboseCommand + " -jar " + jar.getAbsolutePath());
		BufferedReader reader = new BufferedReader(new InputStreamReader(process.getInputStream()));
		while (true) {
			String line = reader.readLine();
			if (line == null) break;
			int start = line.indexOf("[Loaded ");
			if (start < 0) continue;
			if (line.endsWith(jar.getName() + "]")) continue;
			String className = line.substring(8 + start, line.indexOf(" from "));
			action.accept(className);
		}
	}

	private final byte[] bts = new byte[8192];
	private final JarOutputStream jout;
	private final File jar;

	public Compression(File app) throws IOException {
		this.jar = app;
		jout = getJarOutputStream();
	}

	public final void compress() throws IOException {
		List<String> classNames = new LinkedList<>();
		getLoadedClass(jar, classNames::add);
		if (classNames.isEmpty()) return;

		classNames.forEach(this::outputEntry);

		jout.putNextEntry(new JarEntry("META-INF/MANIFEST.MF"));
		jout.write(ManifestInfo.getBytes());
		jout.finish();
		jout.close();
	}

	private void outputEntry(String className) {
		System.err.println("includeing : " + className);
		String resourceName = className.replace('.', '/') + ".class";
		InputStream in = CL.getResourceAsStream(resourceName);
		if (in == null) return;
		int i;
		try {
			jout.putNextEntry(new JarEntry(resourceName));
			while ((i = in.read(bts)) >= 0)
				jout.write(bts, 0, i);
			in.close();
		} catch (IOException e) {
			throw new RuntimeException(e);
		}
	}

	private JarOutputStream getJarOutputStream() throws IOException {
		File file = new File("rt.jar");
		if (!file.exists()) file.createNewFile();
		return new JarOutputStream(new FileOutputStream(file));
	}

	@Deprecated
	public static void getEntriesFromJar(String jarName, Consumer<JarEntry> action) {
		try {
			JarInputStream jin = new JarInputStream(new FileInputStream(jarName));
			while (true) {
				// '.class' -> ''	'/' -> '.'
				JarEntry je = jin.getNextJarEntry();
				if (je == null) break;
				action.accept(je);
			}
			jin.close();
		} catch (IOException e) {
			throw new RuntimeException(e);
		}
	}
}
