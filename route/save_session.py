from .tool.func import *

def save_session():
    data = flask.request.get_json()
    redirect_url = data.get('redirectUrl')

    if redirect_url != '':
        flask.session['redirectUrl'] = redirect_url
    else:
        path = '/' + redirect_url.split('://')[-1].split('/', 1)[-1]
        flask.session['redirectUrl'] = path

    return {'message': 'Session value saved'}