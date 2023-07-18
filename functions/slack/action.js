export async function onRequest(context) {
  const form = await context.request.formData();

  const payloadJson = form.get("payload");

  const payload = JSON.parse(payloadJson);

  const teamDomain = payload.team.domain;
	const action = payload.callback_id;
  const user_id = payload.user.id;
	const name = payload.actions[0].name;
	const value = payload.actions[0].value;
  const response_url = payload.response_url;

	console.log(`Request: TeamDomain: ${teamDomain} Action: ${action} Name: ${name} Value: ${value}`);

  const message = {
    response_type: 'in_channel',
    delete_original: true,
    text: `<@${user_id}>: ${value}`,
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
