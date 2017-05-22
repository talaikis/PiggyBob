 # coding=utf-8
from __future__ import unicode_literals

from django.utils.translation import ugettext as T
from django.contrib.auth.models import AbstractUser
from django.db import models


class DateTimes(models.Model):
    id = models.BigAutoField(primary_key=True)
    date_time = models.DateTimeField(verbose_name=T("Date"), unique=True)

    def __str__(self):
        return u'%s' %(self.nickname)


class CUser(models.Model):
    name = models.CharField(max_length=25, null=True, blank=True)
    nickname = models.CharField(max_length=25, null=True, blank=True)
    email = models.EmailField(max_length=150, verbose_name=T("Email"), primary_key=True)
    location = models.CharField(max_length=255, verbose_name=T("Location"), blank=True, null=True)
    description = models.CharField(max_length=255, verbose_name=T("Description"), blank=True, null=True)

    def __str__(self):
        return u'%s' %(self.nickname)


class UserAuth(models.Model):
    PROVIDERS = (
        (1, "Facebook"),
        (2, "Twitter"),
        (3, "Google"),
        (4, "LinkedIn"),
    )
    user = models.ForeignKey(CUser, on_delete=models.CASCADE, verbose_name=T("User"))
    access_token = models.TextField()
    provider = models.SmallIntegerField(choices=PROVIDERS)
    avatar_url = models.URLField(verbose_name=T("Avatar URL"), null=True, blank=True)
    user_pid = models.IntegerField(verbose_name=T("User ID at provider's"))

    class Meta:
        unique_together = ("user", "provider")


class IncomeCategory(models.Model):
    title = models.CharField(max_length=120, verbose_name=T("Income category"), primary_key=True)
    description = models.TextField(blank=True, null=True, verbose_name=T("Category description"))
    parent = models.ForeignKey('self', verbose_name=T("Parent category"), default=None, blank=True, null=True)

    def __str__(self):
        return u'%s' %(self.title)


class ExpenseCategory(models.Model):
    title = models.CharField(max_length=120, verbose_name=T("Expense category"), primary_key=True)
    description = models.TextField(blank=True, null=True, verbose_name=T("Category description"))
    parent = models.ForeignKey('self', verbose_name=T("Parent category"), default=None, blank=True, null=True)

    def __str__(self):
        return u'%s' %(self.title)


class Currencies(models.Model):
    CURRENCIES = (
        ('AUD', T('Australian dollar')),
        ('CAD', T('Canadian dollar')),
        ('CHF', T('Swiss franc')),
        ('CNY', T('Chinese yuan')),
        ('CZK', T('Czech koruna')),
        ('DKK', T('Danish krone')),
        ('EUR', T('Euro')),
        ('HKD', T('Hong Kong dollar')),
        ('HIF', T('Hungarian forint')),
        ('ISK', T('Icelandic króna')),
        ('JPY', T('Japanese yen')),
        ('MXN', T('Mexican peso')),
        ('NZD', T('New Zealand dollar')),
        ('PLN', T('Polish złoty')),
        ('RUB', T('Russian ruble')),
        ('SEK', T('Swedish krona/kronor')),
        ('SGD', T('Singapore dollar')),
        ('TRY', T('Turkish lira')),
        ('USD', T('United States dollar')),
    )

    title = models.CharField(max_length=3, verbose_name=T("Currency"), primary_key=True, choices=CURRENCIES, default='USD')

    def __str__(self):
        return u'%s' %(self.title)


class Accounts(models.Model):
    id = models.BigAutoField(primary_key=True)
    user = models.ForeignKey(CUser, on_delete=models.CASCADE, verbose_name=T("User"))
    title = models.CharField(max_length=120, verbose_name=T("Account name"))
    currency = models.ForeignKey(Currencies, verbose_name=T("Currency"))

    class Meta:
        unique_together = ("user", "title", "currency")

    def __str__(self):
        return u'%s' %(self.title)


class Income(models.Model):
    id = models.BigAutoField(primary_key=True)
    account = models.ForeignKey(Accounts, on_delete=models.CASCADE, verbose_name=T("Account"))
    category = models.ForeignKey(IncomeCategory, on_delete=models.CASCADE, verbose_name=T("Income Category"))
    date_time = models.ForeignKey(DateTimes, verbose_name=T("Income date"))
    title = models.CharField(max_length=120, verbose_name=T("Income"))
    reference = models.TextField(blank=True, null=True, verbose_name=T("Income description"))
    value = models.FloatField(verbose_name=T("Value"), blank=True, null=True)
    currency = models.ForeignKey(Currencies, verbose_name=T("Currency"))
    rate = models.FloatField(verbose_name=T("Currency rate"), blank=True, null=True)

    class Meta:
        unique_together = ("category", "date_time", "title")

    def __str__(self):
        return u'%s' %(self.title)


class Expense(models.Model):
    id = models.BigAutoField(primary_key=True)
    account = models.ForeignKey(Accounts, on_delete=models.CASCADE, verbose_name=T("User"))
    category = models.ForeignKey(ExpenseCategory, on_delete=models.CASCADE, verbose_name=T("Expense Category"))
    date_time = models.ForeignKey(DateTimes, verbose_name=T("Expense date"))
    title = models.CharField(max_length=120, verbose_name=T("Expense"))
    reference = models.TextField(blank=True, null=True, verbose_name=T("Expense description"))
    value = models.FloatField(verbose_name=T("Value"), blank=True, null=True)
    currency = models.ForeignKey(Currencies, verbose_name=T("Currency"))
    rate = models.FloatField(verbose_name=T("Currency rate"), blank=True, null=True)

    class Meta:
        unique_together = ("category", "date_time", "title")

    def __str__(self):
        return u'%s' %(self.title)
