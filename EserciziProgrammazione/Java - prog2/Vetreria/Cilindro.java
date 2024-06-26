public class Cilindro extends Contenitore {

    private final double altezza;
    private final double raggio;

    public Cilindro(double contenuto, String liquido, double capienza, double altezza, double raggio)
            throws IllegalArgumentException, ExceededCapacityException {
        super(contenuto, liquido, altezza * Math.PI * raggio * raggio);
        if (altezza <= 0) {
            throw new IllegalArgumentException("altezza negativa");
        }
        if (raggio <= 0) {
            throw new IllegalArgumentException("raggio negativo");
        }
        this.altezza = altezza;
        this.raggio = raggio;
    }

}
