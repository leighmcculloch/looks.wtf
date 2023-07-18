import looks from '../../../data/looks.json';

export async function onRequest(context) {
  var tags = new Set();
  for (const look of looks) {
    const lookTags = look.tags.split(/ +/);
    for (const tag of lookTags) {
        tags.add(tag);
    }
  }

  const response = {
    tags: Array.from(tags).toSorted().map((tag) => {
      return { tag, path: `/${tag}`};
    }),
  };

  return new Response(JSON.stringify(response, null, 2), {
    headers: {
      'Content-Type': 'application/json;charset=UTF-8'
    }
  });
}
