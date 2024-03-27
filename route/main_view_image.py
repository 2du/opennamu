from .tool.func import *

def main_view_image(name = ''):
    with get_db_connect() as conn:
        if acl_check(conn, '', 'render') != 1:
            name = re.sub(r'\.cache_v(?:[0-9]+)$', '', name)
            mime_type = re.search(r'([^.]+)$', name)
            if mime_type:
                mime_type = mime_type.group(1).lower()
                if mime_type == 'svg':
                    mime_type = 'svg+xml'

                return flask.send_from_directory('./' + load_image_url(conn), name, mimetype = 'image/' + mime_type)
            else:
                return ''
        else:
            return re_error(conn, '/ban')