const basePath = '/api/v1/'

export default function* ({ tags, looks }) {
  for (const t of tags) {
    yield {
      url: `${basePath}${t}.json`,
      content: JSON.stringify(
          {
            tag: t,
            looks: looks
              .filter((l) => l.tags.includes(t))
              .map((l) => ({ ...l, tags: l.tags.split(/ +/) }))
        },
        null,
        2,
      )
    };
  }
}
