from django.db import models
from django.utils.translation import gettext_lazy as _

# Create your models here.

class User(models.Model):
    id = models.BigAutoField(primary_key=True)

    username = models.CharField(_('username'), max_length=60, unique=True)
    password = models.CharField(_('password'), max_length=255)
    
    email = models.EmailField(
        verbose_name=_('email address'), max_length=255, unique=True
    )

    name = models.CharField(_('name'), max_length=30, blank=True, null=True)
    photo = models.ImageField(upload_to='user_photo/%Y/%m/%D/', blank=True, null=True)  
    birthday = models.DateTimeField(_('club'), blank=True, null=True)
    relationship = models.CharField(_('interest'), max_length=255, blank=True, null=True)
    interest = models.TextField(_('interest'), blank=True, null=True)
    club = models.CharField(_('club'), max_length=255, blank=True, null=True)
    favorite_course = models.TextField(_('favorite_course'), blank=True, null=True)
    favorite_country = models.TextField(_('favorite_country'), blank=True, null=True)
    trouble = models.TextField(_('trouble'), blank=True, null=True)
    exchange = models.TextField(_('exchange'), blank=True, null=True)
    trying = models.TextField(_('trying'), blank=True, null=True)

    friends = models.ManyToManyField("qcard.friend", related_name="friend", blank=True, db_table="user_friends")
    pairing = models.ForeignKey("qcard.friend", related_name="pairing", blank=True, null=True, on_delete=models.SET_NULL)

    class Meta:
        db_table = 'users'

class Friend(models.Model):
    id = models.BigAutoField(primary_key=True)
    user1 = models.ForeignKey('qcard.user', related_name='user1', on_delete=models.CASCADE)
    user2 = models.ForeignKey('qcard.user', related_name='user2', on_delete=models.CASCADE)

    pair = models.BooleanField(default=False)

    class Meta:
        db_table = 'friends'