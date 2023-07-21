export const layout = './_category.vto';

export default function* ({ tags, looks }, filters) {
  for (const tag of tags) {
    yield {
      url: `${filters.tagurl(tag)}index.html`,
      tag,
      tags,
      looks: looks.filter((l) => l.tags.includes(tag))
    };
  }
}
