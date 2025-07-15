package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Article: Representa un artículo en el sistema
type Article struct {
    ID    		primitive.ObjectID 		`json:"id" bson:"_id,omitempty"`
    Title 		string 					`json:"title" bson:"title"`
	Content 	string 					`json:"content" bson:"content"`
	CreatedAt 	time.Time 				`json:"created_at" bson:"created_at"`
}

// Datos de ejemplo para la colección de artículos
var Articles = []Article{
	{
		ID: primitive.NewObjectID(),
		Title: "La Revolución Industrial: El Punto de Inflexión de la Historia Moderna",
		Content: "La Revolución Industrial fue un proceso de transformación económica, social y tecnológica que comenzó en el Reino Unido a mediados del siglo XVIII y se expandió progresivamente por Europa y América del Norte. Este periodo marcó el tránsito de una economía basada en la agricultura y la artesanía a otra dominada por la industria mecanizada, la producción en masa y el capitalismo moderno. Las máquinas de vapor, los telares mecánicos, la minería intensiva y el ferrocarril simbolizaron esta era. Además de sus innovaciones tecnológicas, trajo consigo profundas consecuencias sociales: urbanización acelerada, surgimiento del proletariado industrial, cambios en las estructuras familiares y una creciente desigualdad económica. Fue también un momento crítico para el medio ambiente, pues la explotación de carbón y la contaminación atmosférica comenzaron a tener efectos visibles. La Revolución Industrial no solo cambió la forma en que se producían bienes, sino que transformó la visión del mundo y sentó las bases del sistema global moderno. Sus consecuencias, tanto positivas como negativas, siguen presentes en la sociedad contemporánea.",
		CreatedAt: time.Date(2025, 02, 10, 0, 0, 0, 0, time.UTC),
	},
	{
		ID: primitive.NewObjectID(),
		Title: "Relatividad: La Revolución Conceptual de la Física",
		Content: "La teoría de la relatividad, desarrollada por Albert Einstein a principios del siglo XX, cambió radicalmente nuestra comprensión del espacio, el tiempo y la gravedad. Antes de Einstein, la física clásica newtoniana dominaba el pensamiento científico, considerando el tiempo como absoluto y el espacio como un escenario fijo. Sin embargo, la relatividad especial (1905) introdujo la idea de que el tiempo y el espacio son relativos al observador, y que la velocidad de la luz es constante para todos los marcos de referencia. Diez años después, con la relatividad general (1915), Einstein extendió estos conceptos para incluir la gravedad como una curvatura del espacio-tiempo causada por la masa y la energía. Esta teoría no solo explicó anomalías como la precesión del perihelio de Mercurio, sino que predijo fenómenos nuevos como los agujeros negros, las ondas gravitacionales y la expansión del universo. La relatividad general ha sido confirmada experimentalmente en numerosas ocasiones, incluyendo la reciente detección de ondas gravitacionales. Su impacto trasciende la física: ha influido en la filosofía, en la cosmología y en la forma en que concebimos el universo.",
		CreatedAt: time.Date(2025, 03, 15, 0, 0, 0, 0, time.UTC),
	},
	{
		ID: primitive.NewObjectID(),
		Title: "La Tabla Periódica: Un Mapa del Conocimiento Químico",
		Content: "La tabla periódica de los elementos es uno de los logros más emblemáticos de la ciencia moderna. Organiza todos los elementos químicos conocidos en función de sus propiedades atómicas, permitiendo predecir comportamientos y relaciones entre ellos. Fue creada en 1869 por el químico ruso Dmitri Mendeléyev, quien observó patrones en las masas atómicas y dejó espacios vacíos para elementos aún no descubiertos, muchos de los cuales fueron hallados posteriormente, confirmando su visión. Con el desarrollo de la mecánica cuántica en el siglo XX, la tabla fue reorganizada con base en el número atómico (protones en el núcleo), lo que dio lugar a una estructura aún más coherente. Además de su valor educativo, la tabla periódica es esencial en investigación, medicina, ingeniería y tecnología. De los metales alcalinos a los gases nobles, y de los lantánidos a los elementos sintéticos como el oganesón, cada elemento cuenta una historia de reactividad, estabilidad, estructura electrónica y aplicaciones prácticas. La tabla sigue creciendo con el descubrimiento de nuevos elementos, y representa una síntesis del conocimiento químico humano y su capacidad de ordenar la complejidad natural.",
		CreatedAt: time.Date(2025, 04, 01, 0, 0, 0, 0, time.UTC),
	},
		{
		ID: primitive.NewObjectID(),
		Title: "El ADN: La Clave de la Vida y la Evolución",
		Content: "El ácido desoxirribonucleico (ADN) es la molécula fundamental que codifica la información genética de todos los seres vivos. Su descubrimiento marcó un hito en la biología, al proporcionar una explicación clara de cómo se transmite la herencia. James Watson y Francis Crick, con ayuda de los trabajos de Rosalind Franklin, descubrieron en 1953 la estructura de doble hélice del ADN, un hallazgo que transformó la biología molecular. El ADN está compuesto por cuatro bases nitrogenadas —adenina, timina, guanina y citosina— que se emparejan de forma específica, codificando las instrucciones para construir proteínas, las unidades funcionales de las células. Más allá de la genética clásica, el ADN ha permitido el desarrollo de la ingeniería genética, la biotecnología, la medicina personalizada, y la identificación forense. El Proyecto Genoma Humano, concluido en 2003, secuenció los más de 3 mil millones de pares de bases que componen el genoma humano, abriendo puertas a tratamientos personalizados y nuevas terapias genéticas. En la actualidad, las técnicas como CRISPR permiten editar genes con gran precisión, planteando dilemas éticos sobre el futuro de la biología humana. El ADN no solo es un código, es el libro de la vida, escrito en un alfabeto universal compartido por toda la biodiversidad del planeta.",
		CreatedAt: time.Date(2025, 04, 17, 0, 0, 0, 0, time.UTC),
	},
		{
		ID: primitive.NewObjectID(),
		Title: "Ingeniería Civil: Construyendo las Bases del Mundo Moderno",
		Content: "La ingeniería civil es una de las disciplinas más antiguas y fundamentales de la humanidad. Desde las pirámides de Egipto hasta los rascacielos modernos, esta rama de la ingeniería se encarga del diseño, construcción y mantenimiento de infraestructuras que permiten el desarrollo de la sociedad. Puentes, carreteras, presas, túneles, edificios y sistemas de abastecimiento de agua son ejemplos clásicos. A lo largo de la historia, la ingeniería civil ha evolucionado enormemente, incorporando nuevos materiales como el concreto reforzado, el acero estructural y, más recientemente, los compuestos avanzados. La sostenibilidad se ha convertido en un eje central en la ingeniería contemporánea, impulsando el uso de energías renovables, materiales reciclables y diseño resiliente frente al cambio climático. Hoy, los ingenieros civiles deben combinar conocimientos técnicos con capacidades de gestión de proyectos, normativas medioambientales y urbanismo. Gracias a herramientas como el modelado 3D (BIM) y la simulación estructural, se pueden prever riesgos y optimizar recursos. En definitiva, la ingeniería civil construye no solo estructuras, sino también seguridad, movilidad, desarrollo económico y calidad de vida. Es el cimiento visible e invisible del progreso humano.",
		CreatedAt: time.Date(2025, 06, 24, 0, 0, 0, 0, time.UTC),
	},
}