import java.util.ArrayList;
import java.util.Iterator;

public class Vetreria implements Iterable<Contenitore> {
    ArrayList<Contenitore> vetreria = new ArrayList<>();

    public void add(Contenitore c) {
        vetreria.add(c);
    }

    public void sort() {
        this.vetreria.sort(null);
    }

    public void ottimizza() {
        this.sort();
        for (int i = 0; i < this.vetreria.size(); i++) {
            for (int j = this.vetreria.size() - 1; j > i; j--) {
                this.vetreria.get(j).versa(this.vetreria.get(i));
            }
        }
    }

    @Override
    public Iterator<Contenitore> iterator() {
        return new Iterator<Contenitore>() {

            Iterator<Contenitore> i = vetreria.iterator();

            @Override
            public boolean hasNext() {
                return i.hasNext();
            }

            @Override
            public Contenitore next() {
                return i.next();
            }

        };
    }

}
