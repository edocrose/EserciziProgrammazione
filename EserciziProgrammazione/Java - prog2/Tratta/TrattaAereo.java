public class TrattaAereo extends Tratta {

    private double co2km;
    private double co2decollo;

    public TrattaAereo(String origine, String destinazione, double lung, double vel, double co2decollo, double co2km)
            throws IllegalArgumentException {
        // MODIFIES: this
        // EFFECTS: inizializza una tratta in aereo con gli attributi di Tratta + altri
        // due
        // se i nuovi parametri non sono corretti lancia IllegalArgumentException

        super(origine, destinazione, lung, vel);
        if (co2decollo <= 0) {
            throw new IllegalArgumentException("CO2 decollo <= 0");
        }
        if (co2km <= 0) {
            throw new IllegalArgumentException("CO2 km <= 0");
        }
        this.co2decollo = co2decollo;
        this.co2km = co2km;
    }

    @Override
    public double calcolaCO2() {
        // EFFECTS: restituisce CO2 prodotta come co2decollo + lunghezza * co2km
        return this.co2decollo + (this.getLung() * this.co2km);
    }

    @Override
    public String toString() {
        return "(Aereo) Tratta da " + this.getOrigine() + " a " + this.getDestinazione() + "; lunghezza: "
                + this.getLung() + ", velocità: " + this.getVel() + ", CO2 Dec: " + this.co2decollo + ", CO2: "
                + this.co2km;
    }

}
