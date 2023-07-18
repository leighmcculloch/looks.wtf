import looks from '../../../data/looks.json';

export async function onRequest(context) {
  const tag = context.params.tag;

  console.log(`Request: Tag: ${tag}`);

  var looksSelected = [];
  for (const look of looks) {
    if (look.tags.includes(tag)) {
      looksSelected.push(look);
    }
  }
  if (looksSelected.length == 0) {
    return new Response(`404 Not Found`, { status: 404 });
  }

  const response = {
    tag,
    looks: looksSelected.map((look) => {
      return {
        plain: look.plain,
        title: look.title || undefined,
        tags: look.tags.split(/ +/),
      };
    }),
  };

  return new Response(JSON.stringify(response, null, 2), {
    headers: {
      'Content-Type': 'application/json;charset=UTF-8'
    }
  });
}
