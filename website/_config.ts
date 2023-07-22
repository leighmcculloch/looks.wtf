import lume from "lume/mod.ts";

const site = lume({ src: './src' });

import vento from "lume/plugins/vento.ts";
import sass from "lume/plugins/sass.ts";
import esbuild from "lume/plugins/esbuild.ts";

site.use(vento());
site.use(sass());
site.use(esbuild());

import tags from "../data/tags.json" assert { type: "json" };
import looks from "../data/looks.json" assert { type: "json" };

site.data("tags", tags);
site.data("looks", looks);

site.filter("upperfirst", (s) => s && s[0].toUpperCase() + s.slice(1));
site.filter("titlecase", (s) => s && s.replace(/\w\S*/g, (word) => word.charAt(0).toUpperCase() + word.slice(1).toLowerCase()));
site.filter("tagurl", (t) => t == 'all' ? '/' : `/${t}/`);

site.copy('static', '');

export default site;

