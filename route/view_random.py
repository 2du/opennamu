from .tool.func import *

from .go_api_w_random import api_w_random

async def view_random():
    with get_db_connect() as conn:
        data = json.loads((await api_w_random()).get_data(as_text = True))["data"]
        
        return redirect(conn, '/w/' + url_pas(data))