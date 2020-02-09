# -*- coding: utf-8 -*-
"""
Created on Sat Feb  8 14:07:36 2020

@authors: Akshat, Krishan
"""

import bs4 as bs
import urllib.request
import re
import nltk
import math

from nltk.corpus import stopwords
from gensim.models import Word2Vec
from nltk.tokenize import word_tokenize


class TopicFinder:
    def __init__(self):
        self.dataText = ""
        self.convDict = {0 : "brexit.model",
                         1 : "uselection.model",
                         2 : "coronavirus.model"}
        self.word2vec = Word2Vec.load("brexit.model")
        self.stopWords = stopwords.words('english')

    def loadArticle(self, url):
        scrapped_data = urllib.request.urlopen(url)
        article = scrapped_data.read()
        parsed_article = bs.BeautifulSoup(article, 'lxml')
        paragraphs = parsed_article.find_all('p')
        article_text = ""

        for p in paragraphs:
            article_text += p.text

        # Cleaning the text
        processed_article = article_text.lower()
        # processed_article = re.sub('[^a-zA-Z]', ' ', processed_article)
        # processed_article = re.sub(r'\s+', ' ', processed_article)

        self.dataText = processed_article

    def train(self, convID):
        # Preparing the dataset
        all_sentences = nltk.sent_tokenize(self.dataText)

        all_words = [nltk.word_tokenize(sent) for sent in all_sentences]

        # Removing Stop Words
        for i in range(len(all_words)):
            all_words[i] = [w for w in all_words[i] if w not in self.stopWords]

        self.word2vec = Word2Vec(all_words, min_count=1)
        self.word2vec.save(self.convDict.get(convID))

    def checkInData(self, word):
        try:
            self.word2vec.wv.similar_by_word(word)
            return True
        except KeyError:
            return False

    def checkRelevant(self, string, convID):
        self.word2vec = Word2Vec.load(self.convDict.get(convID))
        words = self.removeStopWords(string.lower())
        print("These are the words", words)
        total = 0
        for word in words:
            # print(self.word2vec.wv.most_similar(word)[0][1])
            if self.checkInData(word):
                print(word)
                print(self.word2vec.wv.most_similar(word))
                total += self.word2vec.wv.most_similar(word)[0][1]
        if total > 0.1:
            return 1
        return 0

    def removeStopWords(self, string):
        word_tokens = word_tokenize(string)
        filtered_sentence = [w for w in word_tokens if not w in self.stopWords]
        return filtered_sentence

if __name__ == "__main__":
    topicFinder = TopicFinder()

    topicFinder.loadArticle('https://en.wikipedia.org/wiki/Brexit')
    topicFinder.loadArticle("https://www.bbc.co.uk/news/world-us-canada-51070020")
    topicFinder.loadArticle("https://www.dailymail.co.uk/news/article-7980883/Video-shows-officials-protective-suits-dragging-suspected-coronavirus-carriers-homes.html")
    topicFinder.train(2)
    print(topicFinder.checkRelevant("What did voters in England and Wales do", 0))
    print(topicFinder.checkRelevant("What is different about next years election", 1))
    print(topicFinder.checkRelevant("There are circuit boards sold in Wuhan China", 2))
