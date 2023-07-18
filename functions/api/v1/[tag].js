import looks from '../../../data/looks.json';

export async function onRequest(context) {
  const tag = context.params.tag;

  console.log(`Request: Tag: ${tag}`);

  var looksWeighted = [];
  for (const look of looks) {
    if (look.tags.includes(tag)) {
      const weight = Math.random();
      looksWeighted.push({ look, weight });
    }
  }
  if (looksWeighted.length == 0) {
    return new Response(`404 Not Found`, { status: 404 });
  }
  looksWeighted.sort((a, b) => a.weight - b.weight);
  const looksSelected = looksWeighted.map((lookWeighted) => lookWeighted.look);

  const response = {
    looks: looksSelected.map((look) => { return { plain: look.plain }; }),
  };

  return new Response(JSON.stringify(response, null, 2), {
    headers: {
      'Content-Type': 'application/json;charset=UTF-8'
    }
  });
}
