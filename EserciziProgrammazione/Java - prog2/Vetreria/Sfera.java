public class Sfera extends Contenitore {

    private double r;

    public Sfera(double contenuto, String liquido, double capienza, double r)
            throws IllegalArgumentException, ExceededCapacityException {
        super(contenuto, liquido, Math.PI * r * r * r * 4 / 3);
        if (this.r <= 0) {
            throw new IllegalArgumentException("Raggio negativo");
        }
        this.r = r;
    }

}
