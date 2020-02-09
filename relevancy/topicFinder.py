# -*- coding: utf-8 -*-
"""
Created on Sat Feb  8 14:07:36 2020

@authors: Akshat, Krishan
"""

import bs4 as bs
import urllib.request
import re
import nltk
import pickle

from nltk.corpus import stopwords
from gensim.models import Word2Vec
from nltk.tokenize import word_tokenize

from os import listdir, getcwd

print(listdir(getcwd()), getcwd())
class TopicFinder:
    def __init__(self):
        self.dataText = ""
        self.topicsDict = {"brexit":  "brexit.pickle",
                           "uselection" : "uselection.pickle",
                           "coronavirus" : "coronavirus.pickle"}
        self.convDict = {"brexit" : "brexit.model",
                         "uselection" : "uselection.model",
                         "coronavirus" : "coronavirus.model"}
        # self.word2vec = Word2Vec.load("brexit.model")
        self.stopWords = stopwords.words('english')

    def loadFile(self, filename):
        file = open(filename, "rb")
        return pickle.load(file)

    def loadArticle(self, url, convID):
        scrapped_data = urllib.request.urlopen(url)
        article = scrapped_data.read()
        parsed_article = bs.BeautifulSoup(article, 'lxml')
        paragraphs = parsed_article.find_all('p')
        article_text = ""

        for p in paragraphs:
            article_text += p.text

        # Cleaning the text
        processed_article = article_text.lower()
        print("THIS IS IT ", len(self.loadFile(self.topicsDict.get(convID))))
        self.dataText = self.loadFile(self.topicsDict.get(convID)) + processed_article
        print("AFTER ", len(self.dataText))
        # self.dataText = article_text
        file = open(convID.split('/')[0] + '.pickle', "wb")
        pickle.dump(self.dataText, file)
        file.close()

        self.dataText = processed_article

    def train(self, convID):
        # Preparing the dataset
        all_sentences = nltk.sent_tokenize(self.dataText)

        all_words = [nltk.word_tokenize(sent) for sent in all_sentences]

        # Removing Stop Words
        for i in range(len(all_words)):
            all_words[i] = [w for w in all_words[i] if w not in self.stopWords]

        self.word2vec = Word2Vec(all_words, min_count=1)
        self.word2vec.save(convID.split('/')[0] + ".model")

    def checkInData(self, word):
        try:
            self.word2vec.wv.similar_by_word(word)
            return True
        except KeyError:
            return False

    def checkRelevant(self, string, convID):
        self.word2vec = Word2Vec.load(convID.split('/')[0] + ".model")
        words = self.removeStopWords(string.lower())
        print("These are the words", words)
        total = 0
        for word in words:
            if self.checkInData(word):
                print(word)
                print(self.word2vec.wv.most_similar(word))
                total += self.word2vec.wv.most_similar(word)[0][1]
        if total > 0.4:
            return 1
        return 0

    def removeStopWords(self, string):
        word_tokens = word_tokenize(string)
        filtered_sentence = [w for w in word_tokens if not w in self.stopWords]
        return filtered_sentence

if __name__ == "__main__":
    topicFinder = TopicFinder()

    # topicFinder.loadArticle('https://www.theguardian.com/commentisfree/2016/feb/22/my-instinct-brexit-boris-anti-eu-not-anti-europe', 'brexit')
    # topicFinder.loadArticle('https://www.nytimes.com/interactive/2019/world/europe/what-is-brexit.html', 'brexit')
    # topicFinder.loadArticle('https://www.thenation.com/article/archive/brexit-eu-england-trade/', 'brexit')
    # topicFinder.loadArticle('https://en.wikipedia.org/wiki/Brexit', 'brexit')
    # topicFinder.loadArticle('https://www.theweek.co.uk/brexit-0', 'brexit')
    # topicFinder.loadArticle('https://www.telegraph.co.uk/brexit/', 'brexit')
    # topicFinder.loadArticle('https://www.thetimes.co.uk/topic/brexit?page=1', 'brexit')

    # topicFinder.train('brexit')

    # topicFinder.loadArticle("https://www.theguardian.com/us-news/us-elections-2020", 'uselection')
    # topicFinder.loadArticle("https://www.reuters.com/politics/us-election2020", 'uselection')
    # topicFinder.loadArticle("https://www.aljazeera.com/news/2020/02/election-2020-presidential-primaries-caucuses-200202201344487.html", 'uselection')
    # topicFinder.loadArticle("https://www.telegraph.co.uk/news/2020/02/05/us-election-2020-democratic-primary-race-winning-battle-delegates/", 'uselection')
    # topicFinder.loadArticle("https://www.telegraph.co.uk/news/0/us-election-candidates-who-democrats-running-2020-presidency-bernie-sanders/", 'uselection')
    # topicFinder.loadArticle("https://www.abc.net.au/news/2020-01-14/when-to-pay-attention-to-us-politics-in-2020/11847092", 'uselection')
    # topicFinder.loadArticle("https://edition.cnn.com/election/2020/primaries-and-caucuses", 'uselection')
    # topicFinder.train('uselection')

    # topicFinder.loadArticle("https://www.aljazeera.com/news/2020/01/coronavirus-symptoms-vaccines-risks-200122194509687.html", 'coronavirus')
    # topicFinder.loadArticle("https://www.businessinsider.com/medical-surveillance-allowed-china-to-discover-novel-coronavirus-2020-1?r=US&IR=T", 'coronavirus')
    # topicFinder.loadArticle("https://www.wired.com/story/wuhan-coronavirus-snake-flu-theory/", 'coronavirus')
    # topicFinder.loadArticle("https://www.vox.com/2020/2/7/21124157/coronavirus-hong-kong-protests-china-carrie-lam", 'coronavirus')
    # topicFinder.loadArticle("https://www.theguardian.com/business/2020/feb/05/coronavirus-threatens-australian-economy-reeling-from-drought-and-fires", 'coronavirus')
    # topicFinder.loadArticle("https://www.nytimes.com/2020/02/08/world/asia/coronavirus-china.html", 'coronavirus')
    # topicFinder.loadArticle("https://www.nytimes.com/2020/02/06/world/asia/coronavirus-china.html", 'coronavirus')
    #
    # topicFinder.train('coronavirus')