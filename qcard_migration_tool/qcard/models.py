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
    gender = models.CharField(_('gender'), max_length=1, blank=True, null=True)
    photo = models.ImageField(
        upload_to='user_photo/%Y/%m/%D/', blank=True, null=True)
    birthday = models.DateTimeField(_('club'), blank=True, null=True)

    relationship = models.CharField(
        _('interest'), max_length=255, blank=True, null=True)

    interest = models.TextField(_('interest'), blank=True, null=True)
    club = models.CharField(_('club'), max_length=255, blank=True, null=True)

    favorite_course = models.TextField(
        _('favorite_course'), blank=True, null=True)
    favorite_country = models.TextField(
        _('favorite_country'), blank=True, null=True)

    trouble = models.TextField(_('trouble'), blank=True, null=True)
    exchange = models.TextField(_('exchange'), blank=True, null=True)
    trying = models.TextField(_('trying'), blank=True, null=True)

    pairing = models.ForeignKey(
        "qcard.friend", related_name="pairing", blank=True, null=True, on_delete=models.SET_NULL)

    class Meta:
        db_table = 'users'


class Friend(models.Model):
    id = models.BigAutoField(primary_key=True)
    user_one = models.ForeignKey(
        'qcard.user', related_name='user1', on_delete=models.CASCADE)
    user_two = models.ForeignKey(
        'qcard.user', related_name='user2', on_delete=models.CASCADE)

    pair = models.BooleanField(default=False)
    created_at = models.DateTimeField(auto_now=True)

    class Meta:
        db_table = 'friends'


class Category(models.Model):
    id = models.BigAutoField(primary_key=True)
    name = models.CharField(max_length=255, unique=True)
    photo = models.ImageField(
        upload_to='cate_photo/%Y/%m/%D/', blank=True, null=True)
    rule = models.TextField(blank=True)
    created_at = models.DateTimeField(auto_now=True)

    class Meta:
        db_table = 'categories'


class Post(models.Model):
    id = models.BigAutoField(primary_key=True)
    creator = models.ForeignKey(User, on_delete=models.CASCADE)
    category = models.ForeignKey(Category, on_delete=models.CASCADE)
    description = models.TextField()
    like = models.ManyToManyField(
        User, related_name="like_post", blank=True, db_table="post_likes")
    created_at = models.DateTimeField(auto_now=True)

    class Meta:
        db_table = 'posts'


class Reply(models.Model):
    id = models.BigAutoField(primary_key=True)
    current_post = models.ForeignKey(Post, on_delete=models.CASCADE)
    creator = models.ForeignKey(User, on_delete=models.CASCADE)
    description = models.TextField()
    like = models.ManyToManyField(
        User, related_name="like_reply", blank=True, db_table="reply_likes")
    created_at = models.DateTimeField(auto_now=True)

    class Meta:
        db_table = 'replys'
