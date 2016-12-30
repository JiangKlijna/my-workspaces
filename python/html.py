
Templates = ['<%s />', '<%s%s />', '<%s>%s</%s>', '<%s%s>%s</%s>', ' %s="%s"']

class Tag(object):

    """docstring for Tag."""
    def __init__(self, name):
        self.name = name
        self.attrs = {}
        self.subs = []
        self.value = None

    def attr(self, *args):
        count = len(args)
        if count % 2 != 0:
            raise Exception('%s = ?' % (args[-1], ))
        for i in range(0, count, 2):
            self.attrs[args[i]] = args[i + 1]
        return self

    def sub(self, *tags):
        if self.value != None:
            raise Exception('self.value = %s' % (self.value, ))
        self.subs += tags
        return self

    def val(self, v):
        if len(self.subs) != 0:
            raise Exception('self.subs = %s' % (self.subs, ))
        self.value = v
        return self

    def build(self, tabnum=-1, isFormat=True):
        if tabnum == -1 and isFormat==True: # -1代表首次执行build
            tabnum = 0

        subs_count, attrs_count = len(self.subs), len(self.attrs)
        if self.value == None and subs_count == 0: #没有值
            if attrs_count == 0: #没有属性
                self.html = Templates[0] % (self.name)
            else: #有属性
                self.html = Templates[1] % (self.name, self.attr_html)
        else:
            if isFormat and self.value == None:
                value_html = self.value_html(tabnum + 1) + '\n' + ('\t' * tabnum)
            else:
                value_html = self.value_html(0)
            if attrs_count == 0: #没有属性
                self.html = Templates[2] % (self.name, value_html, self.name)
            else: #有属性
                self.html = Templates[3] % (self.name, self.attr_html, value_html, self.name)
        if isFormat:
            self.html = '\n' + ('\t' * tabnum) + self.html
        return self

    def save(self, fn):
        f = open(fn, 'w')
        f.write(self.html)
        f.close()

    def value_html(self, tabnum=0):
        if self.value == None:
            cache = []
            for sub in self.subs:
                cache.append(sub.build(tabnum=tabnum, isFormat=(tabnum!=0)).html)
            self.value = ''.join(cache)
        return self.value

    @property
    def attr_html(self):
        cache = []
        for k in self.attrs:
            cache.append(Templates[4] % (k, self.attrs[k]))
        return ''.join(cache)

def main():
    Tag('html').attr('width', '100%').sub(
        Tag('head').sub(
            Tag('meta').attr('name', 'keywords', 'content', 'xxx'),
            Tag('meta').attr('name', 'renderer', 'content', 'webkit'),
            Tag('meta').attr('charset', "utf-8"),
            Tag('title').val('biao ti')
        ),
        Tag('body').sub(
            Tag('br'),
            Tag('div').attr('class', 'style').sub(
                Tag('h2').val('h2'),
                Tag('br'),
                Tag('h3').val('h3')
            ),
            Tag('ul').sub(
                Tag('li').val('1'),
                Tag('li').val('2'),
                Tag('li').val('3'),
                Tag('li').val('4')
            )
        )
    ).build(isFormat=True).save('t.html')

main()
