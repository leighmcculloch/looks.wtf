import looks from '../../data/looks.json';

export async function onRequest(context) {
  const form = await context.request.formData();

  const token = form.get("token");
  if (token != context.env.SLACK_VERIFICATION_TOKEN) {
    return new Response(`401 Unauthorized`, { status: 401 })
  }

  const teamDomain = form.get("team_domain");
	const channelName = form.get("channel_name");
	const userId = form.get("user_id");
	const command = form.get("command");
	const tag = form.get("text");
	const response_url = form.get("response_url");

	console.log(`Request: TeamDomain: ${teamDomain} ChannelName: ${channelName} UserID: ${userId} Command: ${command} Text: ${tag}`);

  var chosenLook = undefined;
  var chosenLookWeight = 0;
  for (const look of looks) {
    if (look.tags.includes(tag)) {
      const weight = Math.random();
      if (weight > chosenLookWeight) {
        chosenLook = look;
        chosenLookWeight = weight;
      }
    }
  }

  if (chosenLook === undefined) {
    return new Response(`Try using the /look command with one of these words: all, angry, annoyed, awe, confused, cool, flipping-tables, happy, kissing, lenny, love, sad, swords, the-look.`);
  }

  const message = {
    response_type: 'in_channel',
    replace_original: true,
    text: `<@${userId}>: ${chosenLook.plain}`,
  };

  const response = await fetch(response_url, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(message),
  });
  if (response.status != 200) {
    throw new Error(`Got non-200 error posting message to Slack API.`);
  }

  return new Response('');
}
